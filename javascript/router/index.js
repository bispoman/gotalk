const express = require('express');
const router = express.Router();
const controller = require('../controller');

router.get('/healthcheck', controller.healthCheck);
router.get('/getperson/:id', controller.getPerson);

module.exports = router;