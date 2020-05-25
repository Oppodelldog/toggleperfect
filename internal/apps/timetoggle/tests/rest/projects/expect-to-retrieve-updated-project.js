client.test("Request executed successfully", function () {
    client.assert(response.status === 200, "Response status is not 200");
});

client.test("Response Body contains id", function () {
    client.assert(response.body["id"] === "WEB-12020", "expects id");
});

client.test("Response Body contains description", function () {
    client.assert(response.body["description"] === "My very first project", "expects description");
});

client.test("Response Body contains closed", function () {
    client.assert(response.body["closed"] === true, "expects closed to be true");
});