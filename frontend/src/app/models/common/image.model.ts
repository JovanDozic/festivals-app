export interface GetUploadURIRequest {
  filename: string;
  fileType: string;
}

export interface GetUploadURIResponse {
  imageURL: string;
  uploadURL: string;
}
