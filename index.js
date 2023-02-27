#!/usr/bin/env node
/*
 * Copyright (c) 2020 | RAJKUMAR S (http://rajkumaar.co.in)
 */

'use strict';

const app = require('./api/server');
app.listen(3000, () => console.log('Listening on port 3000!'));
