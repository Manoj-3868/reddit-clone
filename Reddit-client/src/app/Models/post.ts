export class Post {
    _id: number;
    postname: string;
    url: string;
    description: string;
    votecount: number;
    username: string;
    subredditname: string;
    commentcount: number;
    duration?: string;
    upvote: boolean;
    downvote: boolean;
    constructor(id:number,postname:string,url:string,description: string,votecount: number,username: string,subredditname: string,commentcount: number,duration: string,upvot:boolean,downvote: boolean){
        this._id=id
        this.postname=postname
        this.url=url
        this.description=description
        this.votecount=votecount
        this.username=username
        this.subredditname =subredditname
        this.commentcount=commentcount
        this.upvote=upvot
        this.downvote=downvote
        this.duration=duration
    }
}
