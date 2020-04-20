client.test("Request executed successfully", function () {
    client.assert(response.status === 200, "Response status is not 200");
});

client.test("Response Body contains WEB-1 Project", function () {
    let project1 = response.body.projects[0];
    client.assert(project1.id === "WEB-1", "expects id");
    client.assert(project1.description === "My first project", "expects description");
});

client.test("Response Body contains WEB-2 Project", function () {
    let project2 = response.body.projects[1];
    client.assert(project1.id === "WEB-2", "expects id");
    client.assert(project1.description === "My second project", "expects description");
});