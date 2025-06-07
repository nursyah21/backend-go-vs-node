import test, { expect } from "@playwright/test";

const BASE_URL = 'http://localhost:5000'

test('get /health', async ({ request }) => {
  const newIssue = await request.get(`${BASE_URL}/api/v1/health`);
  expect(newIssue.ok()).toBeTruthy();
});