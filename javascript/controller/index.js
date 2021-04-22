const service = require('../service');

    exports.healthCheck = (req, res) => {
        res.status(200).send("OK");
    };

    exports.getPerson = async (req, res) => {
        const id = req.params.id;
        const response = await service.getPerson(id);
        res.status(200).send(response.data);
    };