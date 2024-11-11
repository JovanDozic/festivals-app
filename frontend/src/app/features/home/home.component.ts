import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../core/auth.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
})
export class HomeComponent implements OnInit {
  isLoggedInStatus: boolean = false;

  constructor(private authService: AuthService) {}

  ngOnInit() {
    this.isLoggedInStatus = this.authService.isLoggedIn();
  }

  isLoggedIn() {
    return this.isLoggedInStatus;
  }
}
