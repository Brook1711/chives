from lxml import etree
import requests
import sys
import re

#reload(sys)

#sys.setdefaultencoding('utf-8')

head = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36'
}


def spider(av):
    url = 'http://bilibili.com/video/av' + str(av)
    print(url)
    html = requests.get(url, headers=head)
    selector = etree.HTML(html.text)
    content = selector.xpath("//html")
    for each in content:
        title = each.xpath('//*[@id="viewbox_report"]/h1/span')
        if title:
            print(title[0].text)
            cid_html_1 = each.xpath('//*[@id="link2"]/@value')
            if cid_html_1:
                cid_html = cid_html_1[0]
                cids = re.findall(r'cid=.+&page', cid_html)
                cid = cids[0].replace("cid=", "").replace("&page", "")
                comment_url = 'http://comment.bilibili.com/' + str(304734284) + '.xml'
                print(comment_url)
                comment_text = requests.get(comment_url, headers=head)
                comment_selector = etree.HTML(comment_text.content)
                comment_content = comment_selector.xpath('//i')
                for comment_each in comment_content:
                    comments = comment_each.xpath('//d/text()')
                    if comments:
                        for comment in comments:
                            print(comment)
                            f.writelines(comment + '\n')
            else:
                print('cid not found!')
        else:
            print('video not found!')

if __name__ == '__main__':
    av = 501811924
    spider(av)

