import { View } from "../../common/types/View.js";
import { debounce, setView, throttle } from "../../lib/lib.js";
import { Div } from "../elements/index.js";

/**
 * Creates a new ListView with the given options.
 * @param {{
 *  items: Array,
 *  itemView: View
 * }} options - The options for creating the ListView.
 * @returns {Div} The container element for the ListView.
 */
export class ListView {
  constructor({
    id,
    items = [],
    itemView,
    provider = async () => {},
    providerParams = {},
    style = {},
  } = {}) {
    this.page = 0;

    this.id = id;

    /**
     * @type {Array}
     */
    this.items = items;

    /**
     * @type {View}
     */
    this.itemView = itemView;

    this.provider = provider;

    this.providerParams = providerParams;

    this.refreshing = false;

    /**
     * @type {Div}
     */
    this.listContainer = new Div({
      id: `list-${this.id}-container`,
      className: "list-container",
      style: style,
    });

    this.listContainer.element.addEventListener("scrollend", (event) => {
      if (this.isScrollAtBottom()) {
        this.fetchMoreItems()
      }
    });

    let query = "";

    for (const key in this.providerParams) {
      if (Object.hasOwnProperty.call(this.providerParams, key)) {
        if ((query = !"")) {
          query += "&";
        }
        const value = this.providerParams[key];
        query += `${key}=${value}`;
      }
    }

    this.providerQueries = query;

    this.fetch(query);

    setView(this);
  }

  /**
 * Checks if the scroll is at the bottom of the list container.
 * @returns {boolean} True if the scroll is at the bottom, false otherwise.
 */
  isScrollAtBottom = throttle(() => {
    const elem = this.listContainer.element;
    return elem.scrollHeight - Math.abs(elem.scrollTop) === elem.clientHeight;
  }, 300)

  async fetch() {
    const query =
      this.providerQueries !== ""
        ? this.providerQueries + `&page=${this.page}`
        : `page=${this.page}`;
    await this.provider(query).then((response) => {
      this.listContainer.element.innerHTML = ""
      this.items = []
      this.page = 0
      response = response || []

      // Add the items to the list
      response.forEach((item) => {
        this.addItem(item);
      });
    });
  }

  /**
   * Adds one or more items to the end of the list.
   * @param {...*} items - The items to add to the list.
   */
  addItem(...items) {
    this.items.push(items);
    items.forEach((item) => {
      this.listContainer.addChild(new this.itemView(item));
    });
  }

  /**
   * Adds one or more items to the beginning of the list.
   * @param {...*} items - The items to add to the list.
   */
  prependItem(...items) {
    this.items.unshift(items);
    items.forEach((item) => {
      this.listContainer.prependChild(new this.itemView(item));
    });
  }

  removeItemView(itemView) {
    this.listContainer.removeChild(itemView);
  }

  async refresh() {
    this.refreshing = true
    this.listContainer.element.innerHTML = ""
    this.items = []
    this.page = 0
    await this.fetch().then(() => this.refreshing = false)
  }

  fetchMoreItems = debounce(() => {
    this.page++;
    const query =
      this.providerQueries !== ""
        ? this.providerQueries + `&page=${this.page}`
        : `page=${this.page}`;
    this.provider(query).then((response) => {
      response = response || [];
      response.forEach((item) => {
        this.addItem(item);
      });
    });
  }, 500);
}
