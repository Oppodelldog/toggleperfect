client.test("Request executed successfully", function () {
    client.assert(response.status === 204, "Response status is not 204");
});