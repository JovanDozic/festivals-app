import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatIconModule } from '@angular/material/icon';
import { MatDialogModule } from '@angular/material/dialog';
import { MatDividerModule } from '@angular/material/divider';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatMenuModule } from '@angular/material/menu';
import { MatTableModule } from '@angular/material/table';
import { Log } from '../../../models/common/log.model';
import { LogsService } from '../../../services/user/logs.service';
import { UserService } from '../../../services/user/user.service';
import { AuthService } from '../../../services/auth/auth.service';

@Component({
  selector: 'app-all-logs',
  imports: [
    CommonModule,
    MatButtonModule,
    MatTooltipModule,
    MatIconModule,
    MatDialogModule,
    MatDividerModule,
    MatCardModule,
    MatChipsModule,
    MatMenuModule,
    MatTableModule,
  ],
  templateUrl: './all-logs.component.html',
  styleUrls: ['./all-logs.component.scss', '../../../app.component.scss'],
})
export class AllLogsComponent implements OnInit {
  private logService = inject(LogsService);
  private authService = inject(AuthService);

  userRole: string = '';
  logs: Log[] = [];
  displayedColumns = ['id', 'username', 'createdAt', 'message'];

  ngOnInit() {
    this.userRole = this.authService.getUserRole() ?? 'ADMINISTRATOR';

    this.logService.getLogs().subscribe((logs) => {
      console.log(logs);
      this.logs = logs;
    });
  }
}
