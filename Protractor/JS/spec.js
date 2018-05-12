// 页面功能
function baidu() { 
    this.open = function () {
        browser.waitForAngularEnabled(false);
        return browser.driver.get('https://www.baidu.com');
    };
    this.getSearchInput = function () {
        return $('#kw');
    };
    this.getSubmitBtn = function () {
        return $('#su');
    };
    this.getResults = async function () {
        await browser.wait(ExpectedConditions.presenceOf($('.result.c-container h3')), 5000);
        return $$('.result.c-container h3 a');
    };
}

// 测试用例
describe('测试百度搜索', function () {
    it('测试protractor官网会不会出现在第一个搜索结果中', async function () {
        var page = new baidu();
        await page.open();

        await page.getSearchInput().sendKeys('protractor');
        await page.getSubmitBtn().click();

        var searchResults = await page.getResults();
        var firstResult = await searchResults[0].getText();
        expect(firstResult).toBe('Protractor - end-to-end testing for AngularJS');
    });
});