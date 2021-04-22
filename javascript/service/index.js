const axios = require ('axios');

    exports.getPerson = async (id) => {
        console.log("bateu no service");
        const url = 'https://swapi.dev/api/people/' + id;
        const response = await axios.get(url);
        return response;
    };