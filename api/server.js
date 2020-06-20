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
    const htlgb = $('.htlgb');
    const data = {};
    data['packageID'] = packageID;
    data['version'] = htlgb.eq(6).text();
    data['installs'] = htlgb.eq(5).text();
    data['size'] = htlgb.eq(3).text();
    data['lastUpdated'] = htlgb.eq(1).text();
    data['rating'] = $('.BHMmbe').eq(0).text();
    data['noOfUsersRated'] = $('.EymY4b').eq(0).text().match(/\d+/)[0];
    data['developer'] = htlgb.eq(htlgb.length === 20 ? 17 : 18).text();
    data['lastCached'] = time();
    return data;
}

const updateCacheAndDoTask = (packageID, res = null, type = null) => {
    const crawlerObj = new Crawler();
    crawlerObj.queue({
        uri: getPlayStoreURL(packageID),
        callback: async (error, result, done) => {
            if (!error) {
                const data = scrapeFromHtml(packageID, result.$);
                if (res) {
                    res.json(type ? shieldsResponse(data, type) : data);
                }
                const collection = await getCollection();
                collection.updateOne({packageID: packageID}, {'$set': data}, {upsert: true});
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
    if (data) {
        res.json(data);
        if (shouldUpdateCache(data['lastCached'])) {
            updateCacheAndDoTask(packageID);
        }
    } else {
        updateCacheAndDoTask(packageID, res);
    }
});

app.get('/:type', async (req, res) => {
    const packageID = req.query.id;
    const type = convertType(req.params.type);
    const data = await getAppData(packageID);
    if (data) {
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
