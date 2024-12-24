import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';

// User interface for type safety
export interface User {
  id?: number; // Optional for creating a new user
  user_name: string; // Required user name
  email: string; // Required valid email
}

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private apiUrl = 'http://localhost:8080/users'; // Base URL for backend API

  constructor(private http: HttpClient) {}

  /**
   * Get all users from the backend.
   * @returns Observable<User[]> - Array of users
   */
  getUsers(): Observable<User[]> {
    return this.http.get<User[]>(this.apiUrl).pipe(
      catchError(this.handleError) // Handle errors
    );
  }

  /**
   * Create a new user.
   * @param user - User object to create
   * @returns Observable<User> - Created user
   */
  createUser(user: User): Observable<User> {
    return this.http.post<User>(this.apiUrl, user).pipe(
      catchError(this.handleError) // Handle errors
    );
  }

  /**
   * Update an existing user.
   * @param id - ID of the user to update
   * @param user - Updated user object
   * @returns Observable<User> - Updated user
   */
  updateUser(id: number, user: User): Observable<User> {
    console.log('Updating user from service:', user);
    return this.http.put<User>(`${this.apiUrl}/${id}`, user).pipe(
      catchError(this.handleError) // Handle errors
    );
  }

  /**
   * Delete a user by ID.
   * @param id - ID of the user to delete
   * @returns Observable<void> - Empty response
   */
  deleteUser(id: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`).pipe(
      catchError(this.handleError) // Handle errors
    );
  }

  /**
   * Generic error handler for HTTP requests.
   * @param error - HttpErrorResponse from the API
   * @returns Observable<never> - Error observable to terminate the stream
   */
  private handleError(error: HttpErrorResponse): Observable<never> {
    let errorMessage: string;

    if (error.error instanceof ErrorEvent) {
      // Client-side or network error
      errorMessage = `An error occurred: ${error.error.message}`;
    } else {
      // Backend error
      errorMessage = `Backend returned code ${error.status}, body was: ${error.error.error || error.message}`;
    }

    console.error(errorMessage); // Log the error to the console (can replace with a logger)
    return throwError(() => new Error(errorMessage));
  }
  
  refreshPage() {
    window.location.reload();
  }
}


