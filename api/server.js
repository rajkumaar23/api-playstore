/*
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */

'use strict';

const express = require('express');
const serverless = require('serverless-http');
const app = express();
const showdown = require('showdown');
const chromium = require('chrome-aws-lambda');
const axios = require('axios');
const MongoClient = require('mongodb').MongoClient;

// Constants
const dbURI = process.env.PLAYSTORE_MONGO_URI;
const dbName = 'db';
const collectionName = 'data';
const labelMap = {
    installs: 'Downloads',
    packageID: 'Package ID',
    version: 'Version',
    size: 'App Size',
    lastUpdated: 'Last Updated On',
    rating: 'Rating',
    developer: 'Developer',
    noOfUsersRated: 'No of users rated'
};


// Utility functions
const time = () => Math.floor(new Date().getTime() / 1000);

const getPlayStoreURL = id => `https://play.google.com/store/apps/details?id=${id}`;

const shouldUpdateCache = dateTime => !dateTime || (time() > (dateTime + 3600));

const shieldsResponse = (data, label) => ({
    schemaVersion: 1,
    label: labelMap[label] || "Badge Type",
    message: data[label] || "Invalid"
});

const getCollection = async () => {
    const client = new MongoClient(dbURI, {useUnifiedTopology: true, useNewUrlParser: true});
    await client.connect();
    return client.db(dbName).collection(collectionName);
}

const getAppData = async (packageID) => {
    const collection = await getCollection();
    const data = await collection.findOne({packageID: packageID});
    if (data && Object.keys(data).includes('_id')) delete data['_id'];
    return data;
}

const convertType = (type) => type === 'downloads' ? 'installs' : type === 'package' ? 'packageID' : type;

const scrapeFromHtml = async (packageID) => {
    const browser = await chromium.puppeteer.launch({
        executablePath: await chromium.executablePath,
        args: Array.from(new Set([...chromium.args, '--no-sandbox', '--disable-setuid-sandbox'])),
        defaultViewport: chromium.defaultViewport,
        headless: true,
        ignoreHTTPSErrors: true,
        ignoreDefaultArgs: ['--disable-extensions'],
    });
    let page = await browser.newPage();

    await page.goto(getPlayStoreURL(packageID));
    await page.waitForNetworkIdle();

    const isError = await page.evaluate(() => {
        return !!document.getElementById('error-section');
    });

    if (isError) {
        return null;
    }

    await page.waitForSelector(".VMq4uf");
    await page.click('.VMq4uf button');

    await page.waitForSelector(".sMUprd");

    const data = await page.evaluate(() => {
        const jobs = document.querySelectorAll(".sMUprd");
        let data = {};
        let dataMap = {
            'Version': 'version',
            'Downloads': 'installs',
            'Updated on': 'lastUpdated',
            'Offered by': 'developer'
        };
        jobs.forEach((item, idx) => {
            if (idx < 8) {
                const firstChildText = item.firstChild.textContent;
                if (dataMap[firstChildText]) {
                    if (firstChildText === "Downloads") {
                        data[dataMap[firstChildText]] = item.children[1].textContent.split(" ")[0];
                    } else {
                        data[dataMap[firstChildText]] = item.children[1].textContent;
                    }
                }
            }
        });
        return data;
    });

    await browser.close();
    data['packageID'] = packageID;
    data['lastCached'] = time();
    return data;
}

const updateCacheAndDoTask = async (packageID, res = null, type = null) => {
    const data = await scrapeFromHtml(packageID);
    if (res) {
        if (data) {
            res.json(type ? shieldsResponse(data, type) : data);
        } else {
            res.json({
                error: 'Invalid package ID'
            })
        }
    }
    const collection = await getCollection();
    collection.updateOne({packageID: packageID}, {'$set': data}, {upsert: true});
}


// Route handlers
app.get('/', async (req, res) => {
    const converter = new showdown.Converter();
    const data = await axios.get('https://raw.githubusercontent.com/rajkumaar23/playstore-api/master/README.md', {responseType: 'text'});
    res.send(converter.makeHtml(data.data));
});

app.get('/json', async (req, res) => {
    const packageID = req.query.id;
    const data = await getAppData(packageID);
    if (data && !!req.query.forceUpdate !== true) {
        res.json(data);
        if (shouldUpdateCache(data['lastCached'])) {
            updateCacheAndDoTask(packageID);
        }
    } else {
        updateCacheAndDoTask(packageID, res);
    }
});

app.get('/used-by', async (req, res) => {
    const collection = await getCollection()
    collection.estimatedDocumentCount({}, (error, count) => {
        if (error) {
            res.json({
                error
            })
        }
        res.json({
            schemaVersion: 1,
            label: "Used by",
            message: `${count} apps`
        })
    })
})

app.get('/:type', async (req, res) => {
    const packageID = req.query.id;
    const type = convertType(req.params.type);
    const data = await getAppData(packageID);
    if (data && !!req.query.forceUpdate !== true) {
        res.json(shieldsResponse(data, type));
        if (shouldUpdateCache(data['lastCached'])) {
            updateCacheAndDoTask(packageID);
        }
    } else {
        updateCacheAndDoTask(packageID, res, type);
    }
});

module.exports = app;
module.exports.handler = serverless(app);
