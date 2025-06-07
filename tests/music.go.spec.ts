import test, { expect } from "@playwright/test";

const BASE_URL = 'http://localhost:5000'

test('get /music', async ({ request }) => {
  const newIssue = await request.get(`${BASE_URL}/api/v1/music`);
  expect(newIssue.ok()).toBeTruthy();
});