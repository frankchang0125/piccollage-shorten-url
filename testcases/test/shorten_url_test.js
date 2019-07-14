import {
  expect,
} from 'chai';
import {
  createShortenUrl,
  visitShortenUrlWithoutRedirect,
} from './utils';

/* eslint-disable func-names, no-unused-expressions */
describe('Test shorten url service', function () {
  it('Create shorten url', async () => {
    const url = 'http://www.google.com.tw';
    const shorten = await createShortenUrl(url);
    expect(shorten).not.eq('');
  });

  it('Visit shorten url', async () => {
    const url = 'http://www.facebook.com.tw';
    const shorten = await createShortenUrl(url);
    const okStatus = [301];
    const resp = await visitShortenUrlWithoutRedirect(shorten, okStatus);
    expect(resp.status).to.eq(301);
    expect(resp.header).to.have.property('location', url);
    expect(resp.redirect).to.be.true;
  });

  it('Same urls should have same shorten url', async () => {
    const url = 'http://www.facebook.com.tw';
    const shorten1 = await createShortenUrl(url);
    const shorten2 = await createShortenUrl(url);
    expect(shorten1).to.eq(shorten2);
  });

  it('Visit with non-exists shorten url should generate 404', async () => {
    const shorten = 'KJ2nx8';
    const okStatus = [404];
    const resp = await visitShortenUrlWithoutRedirect(shorten, okStatus);
    expect(resp.status).to.eq(404);
  });
});
/* eslint-enable func-names, no-unused-expressions */
