import requests
import re
from lxml import etree

headers = {"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.87 Safari/537.36"}

class Danmu:
    def __init__(self, appear_time, disappear_time, mode, color, text):
        self.appear_time = appear_time
        self.disappear_time = disappear_time
        self.mode = mode
        self.color = color
        if self.color == r"\c&HFFFFFF": # 默认白色则删去,压缩空间
            self.color = ""
        self.text = text

def get_cid(url):
    html = requests.get(url, headers = headers).content.decode('utf-8')
    html1 = requests.get(url, headers=headers)
    selector = etree.HTML(html1.text)
    content = selector.xpath("//html")
    for each in content:
        title = each.xpath('//*[@id="viewbox_report"]/h1/span')
        if title:
            name = title[0].text
        else:
            print('url wrong')
    if "cid" in html:
        return re.search(r'"cid":(\d*)', html).group(1), name
    else:
        return "x"


def format_time(time_raw):
    hour = str(int(time_raw / 3600))
    if hour != '0':
        time_raw %= 3600
    minute = str(int(time_raw / 60))
    if minute != '0':
        time_raw %= 60
    if len(minute) == 1:
        minute = "0" + minute
    second = str(round(time_raw, 2))
    if '.' in second:
        if len(second.split('.')[0]) == 1:
            second = "0" + second
        if len(second.split('.')[1]) == 1:
            second = second + "0"
    else:
        if len(second) == 1:
            second = "0" + second + ".00"
        else:
            second = second + ".00"
    return hour + ':' + minute + ':' + second

def get_danmu_list(danmu_url):
    danmu_raw = requests.get(danmu_url, headers = headers).content.decode('utf-8')
    danmu_raw_list = re.findall(r"<d .*?</d>", danmu_raw)
    danmu_list = []
    for item in danmu_raw_list:
        m = re.match(r'<d p="(.*?),(.*?),(.*?),(.*?),(.*?)>(.*?)</d>', item) # group(3和5)废弃
        danmu_list.append(Danmu(format_time(float(m.group(1))), format_time(float(m.group(1)) + 8), int(m.group(2)), r"\c&H" + str(hex(int(m.group(4))))[2:].upper(), m.group(6)))
    return danmu_list

def generate_ass(danmu_list):
    currentY = 1 # currentY * 25 -> Y .. currentY = 1 -> 16
    for danmu in danmu_list:
        style = ""
        startX = 620 # likely to be the average
        endX = -(startX - 560)
        if 1 <= danmu.mode and danmu.mode <= 3:
            style = "R2L"
        else:
            style = "FIX"
            startX = 280
            endX = 280
        Y = currentY * 25
        currentY += 1
        if currentY == 17:
            currentY = 1
        ass += "Dialogue: 0,%s,%s,%s,,20,20,2,,{\\move(%d,%d,%d,%d)%s}%s\n" % (danmu.appear_time, danmu.disappear_time, style, startX, Y, endX, Y, danmu.color, danmu.text)
    return ass

# https://www.bilibili.com/video/BV1PJ411r7Sd

def main():
    url = input("url:")
    cid, name = get_cid(url)
    if cid == "x":
        print("cid not found")
        exit(1)
    danmu_url = "https://comment.bilibili.com/" + cid + ".xml"
    danmu_list = get_danmu_list(danmu_url)
    ass = generate_ass(danmu_list)
    with open(name+ ".ass", "wb+") as ass_file:
        ass_file.write(ass.encode("utf-8"))
    # print(name)

if __name__ == "__main__":
    main()