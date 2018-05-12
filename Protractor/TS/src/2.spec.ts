import { Baidu } from './page';

describe('测试百度搜索', function () {
    it('测试protractor官网会不会出现在第一个搜索结果中', async function () {
        let baidu = new Baidu();
        await baidu.open();

        await baidu.getSeachInput().sendKeys('protractor');
        await baidu.getSubmitBtn().click();

        let results = await baidu.getResults();
        let firstResult = await results[0].getText();
        expect(firstResult).toBe('Protractor - end-to-end testing for AngularJS');
    });
});