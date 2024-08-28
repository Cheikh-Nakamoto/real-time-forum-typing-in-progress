import { getViewPager, getPage, getPageByPath } from "../lib.js"

export function goTo (pageId) {
    const page = getPage(pageId)
    const vp = getViewPager(page.pagerId)
    vp.setCurrentPage(page)
}


export function router(){
    const page = getPageByPath(location.pathname)
    goTo(page.id)
}