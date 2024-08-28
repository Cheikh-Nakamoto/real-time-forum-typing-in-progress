export class Comment {
  constructor(data) {
    this._author = data.username || data.email;
    this._avatar = data.avatar;
    this._duration = data.duration;
    this._content = data.content;
    this._likes = data.likes;
  }
  get author() {
    return this._author;
  }
  get avatar() {
    return this._avatar;
  }
  get duration() {
    return this._duration;
  }
  get content() {
    return this._content;
  }
  get likes() {
    return this._likes;
  }
  get dislikes() {
    return this._dislikes;
  }
}
