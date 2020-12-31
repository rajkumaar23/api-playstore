/*
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */

'use strict';

const express = require('express');
const serverless = require('serverless-http');
const app = express();
const showdown = require('showdown');
const Crawler = require("crawler");
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

const scrapeFromHtml = (packageID, $) => {
    const data = {};
    const groupIdentifier = $('.hAyfc');
    let length = groupIdentifier.length;
    let dataMap = {
        'Current Version': 'version',
        'Installs': 'installs',
        'Size': 'size',
        'Updated': 'lastUpdated',
        'Offered By': 'developer'
    };
    while (length--) {
        const element = groupIdentifier.eq(length);
        const firstChildText = element.children(':first-child').text();
        if (dataMap[firstChildText]) {
            data[dataMap[firstChildText]] = element.children(':nth-child(2)').text();
        }
    }
    data['packageID'] = packageID;
    data['rating'] = $('.BHMmbe').eq(0).text();
    data['noOfUsersRated'] = $('.EymY4b').eq(0).text().split(' ')[0];
    data['lastCached'] = time();
    return data;
}

const updateCacheAndDoTask = (packageID, res = null, type = null) => {
    const crawlerObj = new Crawler();
    crawlerObj.queue({
        uri: getPlayStoreURL(packageID),
        callback: async (error, result, done) => {
            if (result.statusCode === 200) {
                const data = scrapeFromHtml(packageID, result.$);
                if (res) {
                    res.json(type ? shieldsResponse(data, type) : data);
                }
                const collection = await getCollection();
                collection.updateOne({packageID: packageID}, {'$set': data}, {upsert: true});
            } else {
                if (res) {
                    res.json({
                        error: 'Invalid package ID'
                    })
                }
            }
            done();
        }
    })
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
