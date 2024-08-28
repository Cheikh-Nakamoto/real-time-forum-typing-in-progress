import { getPage, setViewPager } from "../lib.js"

export class ViewPager {
    constructor(options) {
        this.id = options.id
        this.pages = options.pages.map(page => new page.page({ id: page.id, pagerId: options.id }))
        this.currentPage = options.defaultPage || this.pages[0]
        this.container = options.container
        this.locPathname = options.locPathname

        addEventListener("pageChanged", (event) => {
            const data = event.detail
            if (this.id == data.pagerId) {
                const page = getPage(data.page.id) ||  new data.page({ pagerId: this.id })
                this.container.replaceContent(page)
                window.history.pushState(data.pagerId, page.title, page.path);
                document.title = page.title
            }
        })

        this.setCurrentPage(this.currentPage)
        setViewPager(this)

        addEventListener
    }

    setCurrentPage(page) {
        this.currentPage = page
        const pageChangedEvnt = new CustomEvent("pageChanged", { detail: { pagerId: this.id, page } })
        dispatchEvent(pageChangedEvnt)
    }

}
