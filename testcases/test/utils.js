import superagent from 'superagent';
import logger from './logger';

const host = process.env.HOST;
const port = process.env.PORT;
const httpUrl = `http://${host}:${port}`;

export const createShortenUrl = async (longUrl) => {
  try {
    const body = { url: longUrl };
    const url = `${httpUrl}/shorten`;
    const resp = await superagent.post(url).send(body);
    return resp.body.url;
  } catch (e) {
    logger.error(`Cannot create shorten url: ${e}`);
    throw e;
  }
};

export const visitShortenUrl = async (shorten) => {
  try {
    const url = `${httpUrl}/${shorten}`;
    const resp = await superagent.get(url);
    return resp;
  } catch (e) {
    logger.error(`Cannot visit shorten url: ${e}`);
    throw e;
  }
};

export const visitShortenUrlWithoutRedirect = async (shorten, okStatus = []) => {
  try {
    const url = `${httpUrl}/${shorten}`;
    const resp = await superagent.get(url).redirects(0)
      .ok(res => okStatus.includes(res.status));
    return resp;
  } catch (e) {
    logger.error(`Cannot visit shorten url: ${e}`);
    throw e;
  }
};
