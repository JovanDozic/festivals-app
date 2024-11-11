import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../core/auth.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
})
export class HomeComponent implements OnInit {
  isLoggedInStatus: boolean = false;
  userRole: string | null = null;

  constructor(private authService: AuthService) {}

  ngOnInit() {
    this.isLoggedInStatus = this.authService.isLoggedIn();
    this.userRole = this.authService.getUserRole();
  }

  isLoggedIn() {
    return this.isLoggedInStatus;
  }
}
