export class Comment {
    _id?:number
    text: string;
    postid: number;
    username?:string;
    instant?: string;
    constructor(id:number,text:string,postid:number,user:string,duration:string){
        this._id=id
        this.text=text
        this.postid=postid
        this.username=user
        this.instant=duration
    }
}