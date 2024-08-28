import { setView } from "../../lib/lib.js";

/**
* Creates a new View element.
* @param {{
*   id?: string,
*   className?: string,
*   name?: string,
*   children?: View[],
*   style?: object,
*   placeholder?: string
* }} options - The options for creating the view.
*/
export class View {
  constructor(options = {}) {
    this.element;
    this.children = options.children || [];
    this.style = options.style || {};
    this.classList = options.className && options.className.split(" ") || []
    this.name = options.name;
    this.id = options.id;
    this.placeholder = options.placeholder;
    this.listeners = options.listeners
    this.onScroll = options.onScroll;
    this.constraints = options.constraints
    this.attr = options.attr

    setView(this);
  }

  /**
   * Sets the attributes of the view element.
   */
  _setAttributes() {
    if (this.classList.length > 0)
      this.element.classList.add(...this.classList);
    if (this.id)
      this.element.id = this.id;
    if (this.name)
      this.element.name = this.name;
    if (this.placeholder)
      this.element.placeholder = this.placeholder;

    if (this.attr) {
      for (const attr in this.attr) {
        if (Object.hasOwnProperty.call(this.attr, attr)) {
          const value = this.attr[attr];
          this.element.setAttribute(attr, value);
        }
      }
    }
  }

  _setEventListeners() {
    if (this.listeners) {
      // document.addEventListener("DOMContentLoaded", () => {
      for (const eventName in this.listeners) {
        if (Object.hasOwnProperty.call(this.listeners, eventName)) {
          const func = this.listeners[eventName];
          if (func) {
            if (eventName == 'onclick') {
              this.element.classList.add('clickable')
            }
            this.element[eventName] = func
          }
        }
      }
      // })
    }
  }

  _setConstraints() {
    if (this.constraints) {
      document.addEventListener("DOMContentLoaded", () => {

        if (this.constraints.top) {
          const top = this.constraints.top();
          this.element.style.top = top
        }
        if (this.constraints.bottom) {
          this.element.style.bottom = this.constraints.bottom();
        }
        if (this.constraints.left) {
          this.element.style.left = this.constraints.left();
        }
        if (this.constraints.right) {
          this.element.style.right = this.constraints.right();
        }
      });
    }
  }

  /**
   * Appends the child elements to the view element.
   */
  _appendChildren() {
    this.children.forEach((child) => {
      if (child instanceof Node) {
        this.element.appendChild(child)
      } else {
        this.element.appendChild(child.element)
      }
    });
  }

  /**
   * Applies the styles to the view element.
   */
  _applyStyles() {
    this.style &&
      Object.entries(this.style).forEach(
        ([key, value]) => (this.element.style[key] = value)
      );
  }

  
  /**
   * Adds a child element to the view.
   * @param {View} child - The child element to add.
   */
  addChild(child) {
    this.children.push(child);
    this.element.appendChild(child.element);
  }

  removeChild(child){
    this.children.splice(this.children.indexOf(child), 1);
    const elem = this.element.querySelector(`#${child.id}`);
    this.element.removeChild(elem);
  }

  replaceContent(view) {
    this.children = []
    this.element.innerHTML = ''
    this.addChild(view)
  }

  /**
   * Prepends a child element to the view.
   * @param {View} child - The child element to prepend.
   */
  prependChild(child) {
    this.children.unshift(child);
    this.element.insertBefore(child.element, this.element.firstChild);
  }

  hide() {
    this.element.style.display = "none";
  }

  show() {
    this.element.style.display = "block";
  }
}
