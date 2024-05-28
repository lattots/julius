const axios = require('axios');

// Define the URL for the POST request
const url = 'http://localhost:3000/event';

// Get the current time and calculate the time one-hour later
const startTime = new Date();
const endTime = new Date(startTime.getTime() + 60 * 60 * 1000);

// Format the times to ISO strings
const startISO = startTime.toISOString();
const endISO = endTime.toISOString();

// Define the event data to be sent in the POST request
const event_data = {
    name: 'Sample Event',
    host: 'John Doe',
    location: '123 Event St, City',
    start: startISO,
    end: endISO,
    'dress-code': 'Casual',
    theme: 'Summer Party',
    price: 420.69,
    'signup-link': 'https://example.com/signup'
};

// Send the POST request with the event data as JSON
axios.post(url, event_data)
    .then(response => {
        console.log(`Status Code: ${response.status}`);
        console.log(`Response Text: ${response.data}`);
    })
    .catch(error => {
        console.error(`Error: ${error}`);
    });
