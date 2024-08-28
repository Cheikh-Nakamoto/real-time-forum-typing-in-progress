export class Post {
  constructor(data) {
    this._author = data.username || data.email;
    this._duration = data.Age;
    this._title = data.title;
    this._content = data.body;
    this._categories = data.categories;
    this._commentsOunt = data.CommentsCount;
    this._likes = data.Votes;
    this.userId = data.user_id
    // this._dislikes = data.dislikes;
    this._bookmark = data.bookmark;
    this.userReaction = data.UserReact
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
  get title() {
    return this._title;
  }
  get content() {
    return this._content;
  }
  get categories() {
    return this._categories;
  }
  get comments() {
    return this._comments;
  }
  get likes() {
    return this._likes;
  }
  get dislikes() {
    return this._dislikes;
  }
  get bookmark() {
    return this._bookmark;
  }

  // fromJson (jsonObject) {
  //   this._author = data.username || data.email;
  //   this._avatar = data.avatar;
  //   this._duration = data.age;
  //   this._title = data.title;
  //   this._content = data.body;
  //   this._categories = data.categories;
  //   this._commentsOunt = data.commentsCount;
  //   this._likes = data.likes;
  //   // this._dislikes = data.dislikes;
  //   this._bookmark = data.bookmark;
  //   this.userReaction = data.userReaction
  // }
}
