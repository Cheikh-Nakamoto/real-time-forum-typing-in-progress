export const parseHTML = (htmlStr) => {
    const parser = new DOMParser()
    const document = parser.parseFromString(htmlStr,'text/html')
    return document.body.firstChild
}