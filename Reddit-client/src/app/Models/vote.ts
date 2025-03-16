export class Vote {
    votetype: number;
    postid: number;
    username:string;
    constructor(vote:number,post:number,user:string){
        this.votetype=vote
        this.postid=post
        this.username=user
    }
}
