import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, retry, throwError } from 'rxjs';
import { Post } from '../Models/post';
import { Postrequest } from '../Models/postrequest';

@Injectable({
  providedIn: 'root'
})
export class PostService {
  username:string;
  httpHeader={
    headers:new HttpHeaders({
      'Content-Type':'application/json'
    }
    )
  }
  baseUrl:string="http://localhost:3000"
  constructor(private httpClient:HttpClient) {
    this.username = localStorage.getItem("username")+''
   }

  getAllPosts(username:string): Observable<Array<Post>> {
    console.log('get-all-posts')
    return this.httpClient.get<Array<Post>>(this.baseUrl+'/post-all/'+username).pipe(
      retry(1),
      catchError(this.httpError)
    );
  }

  createPost(postPayload: Postrequest): Observable<any> {
    console.log('create-post')
    return this.httpClient.post(this.baseUrl+'/post',JSON.stringify(postPayload),this.httpHeader).pipe(
      retry(1),
      catchError(this.httpError)
    );
  }

  getPost(id: number): Observable<Post> {
    console.log('get-post')
    return this.httpClient.get<Post>(this.baseUrl+'/post/' + id+'/'+this.username).pipe(
      retry(1),
      catchError(this.httpError)
    );
  }

  getAllPostsByUser(name: string): Observable<Post[]> {
    console.log('get-all-postbyuser')
    return this.httpClient.get<Post[]>(this.baseUrl+'/post/by-user/' + name).pipe(
      retry(1),
      catchError(this.httpError)
    );
  }

  httpError(error:HttpErrorResponse){
    let msg='';
    if(error.error instanceof ErrorEvent){
      msg=error.error.message;
    }
    else{
      msg=`Error Code:${error.status}\nMessafe:${error.message}`;
    }
    console.log(msg);
    return throwError(msg);
  }

}
