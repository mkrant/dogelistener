from bs4 import BeautifulSoup

with open('../backend/static/data/1/index.html', 'r') as file:
    html_doc = file.read()

    soup = BeautifulSoup(html_doc, 'html.parser')
    bodyTag = soup.find('body')
    divTag = soup.new_tag('div')

    prevTag = soup.new_tag('a')
    prevTag['style'] = 'font-size: 30px'
    prevTag['id'] = 'previous'
    prevTag['href'] = 'changeme'
    prevTag.insert(0, 'Previous')

    nextTag = soup.new_tag('a')
    nextTag['style'] = 'font-size: 30px'
    nextTag['id'] = 'next'
    nextTag['href'] = 'changeme'
    nextTag.insert(0, 'Next')

    scriptTag = soup.new_tag('script')
    scriptTag['type'] = "text/javascript"

    js = """
    let test = window.location.pathname.match("\\\\d+") ?? 1;
    console.log(test);
    let num = parseInt(test);
    console.log(num)
    
    if (num-1 > 0) {
        document.getElementById("previous").href = window.location.origin + "/data/" + (num-1);
    } else {
        console.log("lower");
        document.getElementById("previous").hidden = true;
    }
    
    document.getElementById("next").href = window.location.origin + "/data/" + (num+1);
    """
    scriptTag.insert(0, js)

    divTag.append(prevTag)
    divTag.append(nextTag)
    divTag.append(scriptTag)

    bodyTag.insert(0, divTag)

    with open("../backend/static/data/1/output.html", "w") as wfile:
        wfile.write(str(soup))