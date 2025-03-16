import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, retry, throwError } from 'rxjs';
import { Comment } from '../Models/comment';
@Injectable({
  providedIn: 'root'
})
export class CommentService {
  httpHeader={
    headers:new HttpHeaders({
      'Content-Type':'application/json'
    }
    )
  }
  baseUrl:string="http://localhost:3000"
  constructor(private httpClient:HttpClient) { }

  getAllCommentsForPost(postId: number): Observable<Comment[]> {
    return this.httpClient.get<Comment[]>(this.baseUrl+'/comment/by-post/' + postId).pipe(
      retry(1),
      catchError(this.httpError)
    );
  }

  postComment(commentPayload: Comment): Observable<any> {
    return this.httpClient.post<any>(this.baseUrl+'/comment',JSON.stringify(commentPayload),this.httpHeader).pipe(
      retry(1),
      catchError(this.httpError)
    );
  }

  getAllCommentsByUser(name: string): Observable<Comment[]> {
    return this.httpClient.get<Comment[]>(this.baseUrl+'/comment/by-user/' + name).pipe(
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
