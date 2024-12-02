import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../services/auth/auth.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
  standalone: true,
})
export class HomeComponent implements OnInit {
  isLoggedInStatus = false;
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
