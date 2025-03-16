export class Subreddit {
    _id?: number;
    name: string;
    description: string;
    numberOfPosts?: number;
    constructor(id:number,name:string,description:string,nopost:number){
        this._id=id
        this.name=name
        this.description=description
        this.numberOfPosts=nopost
    }
}
