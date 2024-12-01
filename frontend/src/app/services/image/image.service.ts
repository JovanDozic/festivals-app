import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import {
  GetUploadURIRequest,
  GetUploadURIResponse,
} from '../../models/common/image.model';
import { catchError, map, Observable, switchMap } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ImageService {
  private apiUrl = 'http://localhost:4000';

  constructor(private http: HttpClient) {}

  getPresignedURL(request: GetUploadURIRequest) {
    return this.http.post<GetUploadURIResponse>(
      `${this.apiUrl}/image/upload`,
      request
    );
  }

  uploadProfilePhoto(file: File): Observable<any> {
    const formData = new FormData();
    formData.append('file', file);

    const uploadRequest: GetUploadURIRequest = {
      filename: file.name,
      fileType: file.type,
    };

    return this.http
      .post<GetUploadURIResponse>(`${this.apiUrl}/image/upload`, uploadRequest)
      .pipe(
        switchMap((response: GetUploadURIResponse) => {
          const uploadURL = response.uploadURL;
          const imageURL = response.imageURL;

          return this.http
            .put(uploadURL, file, {
              headers: {
                'Content-Type': file.type,
              },
            })
            .pipe(
              map(() => {
                console.log('Profile photo uploaded successfully', imageURL);
                return { imageURL };
              }),
              catchError((error) => {
                console.error('Error uploading profile photo:', error);
                throw error;
              })
            );
        })
      );
  }
}
