const ViewElement = {
    views: new Map(),
    pages: new Map(),
    pagers: new Map(),
    openedChat: new Map(),
}

export function getOpenedChat(chatId){
  return ViewElement.openedChat.get(chatId)
}

export function setOpenedChat(chat){
  ViewElement.openedChat.set(chat.id,chat)
}

export function remOpenedChat(chatId){
  const view =document.getElementById(chatId)
  if (view) {
    view.remove()
  }
  ViewElement.openedChat.delete(chatId)
}

export function resetOpenedChat(){
  ViewElement.openedChat.forEach((chat)=>{
    remOpenedChat(chat.id)
  })
}

export function getView(viewId){
    return ViewElement.views.get(viewId);
}

export function setView(view){
    ViewElement.views.set(view.id,view)
}

export function remView(viewId){
    ViewElement.views.delete(viewId)
}

export function setPage(page){
    ViewElement.pages.set(page.id,page)
}

export function getPage(pagesId){
    return ViewElement.pages.get(pagesId)
}

export function getPageByPath(path){
    const entry =Array.from(ViewElement.pages.entries()).find(([,page])=>page.path === path)
    return entry && entry[1]
}

export function setViewPager(pager){
    ViewElement.pagers.set(pager.id,pager)
}

export function getViewPager(pagerId){
    return ViewElement.pagers.get(pagerId)
}

export function throttle(func, limit) {
  let inThrottle;
  return function() {
    const args = arguments;
    if (!inThrottle) {
        inThrottle = true;
        setTimeout(() => inThrottle = false, limit);
        return func.apply(this, args);
    }
  };
}

export function debounce(func, delay) {
  let timeoutId;
  return function() {
    const context = this;
    const args = arguments;
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => {
      return func.apply(context, args);
    }, delay);
  };
}

export function formatDate(date) {
    const days = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
    const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
  
    const dayName = days[date.getDay()];
    const monthName = months[date.getMonth()];
    const day = date.getDate();
    const year = date.getFullYear();
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');
  
    return `${dayName} ${monthName} ${day} ${year} ${hours}:${minutes}:${seconds}`;
  }