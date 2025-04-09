import type { Response } from '@/stores/models.ts'

const baseURL = '/search/instant/mixed?type=track&page=0&pageSize=5&withLikesCount=true&text=';

export async function getSongs(text: string): Promise<Response> {

  const finalUrl = `${baseURL}${encodeURI(text)}`;

  const response = await fetch(finalUrl, {
    method: 'GET'
  })

  if (!response.ok) {
    console.error(`yandex music search failed, code ${response.status}, data: ${response.text}`);
    throw new Error("yandex music search failed");
  }

  return await response.json() as Response;
}
