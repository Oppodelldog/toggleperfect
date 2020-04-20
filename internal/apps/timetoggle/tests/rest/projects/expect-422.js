client.test("Request executed successfully", function () {
    client.assert(response.status === 422, "Response status is not 422");
});