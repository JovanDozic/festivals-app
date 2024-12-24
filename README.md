# Festivals App

## Overview

Festivals App is a full-stack application designed to streamline the management of music festivals. The app provides a seamless experience for both festival staff and attendees, featuring subsystems for ticket sales, custom package creation (including camp and travel packages), and bracelet management.

## Features

- Comprehensive Event Management: Easily add, update, or delete festival details, including key information such as dates, descriptions, and ticket availability.
- Flexible Ticket Sales: Support a wide range of ticket categories with dynamic pricing options and availability tracking.
- Customizable Packages: Allow attendees to create personalized packages with options for camping, transportation, and additional add-ons to enhance their experience.
- Advanced Wristband Features: Manage the entire lifecycle of festival wristbands, from issuance to usage, including activation and balance top-ups.
- Integrated Tracking Systems: Enable both attendees and organizers to monitor wristband statuses, orders, and access insights to improve logistics and experiences.
- Automated Notifications: Keep users informed with real-time updates on their orders, wristband statuses, and other festival-related activities via email notifications.

## Tech

### Backend:

- **Programming Language:** Go
- **Database:** PostgreSQL
- **Storage:** AWS S3 for image and media storage

### Frontend:

- **Framework:** Angular
- **Design System:** Angular Material with Material 3 guidelines

## Database Design

The database follows a normalized structure to ensure data integrity and efficiency. Key design artifacts include:

- **Entity-Relationship (ER) Diagrams**
- **Relational Models**
- **Sequence Diagrams** (for some features)
- **Use Case Diagrams**

## Setup

### Backend

1. Clone the repository:
   ```bash
   git clone https://github.com/JovanDozic/festivals-app.git
   cd festivals-app/backend
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up environment variables for database and AWS S3 access (see `docs/setup/backend-setup/`)
4. Run the backend server using:
   ```bash
   go run main.go
   ```
   or using debugger in Visual Studio Code (needs `launch.json` from `backend-setup`).

### Frontend

1. Navigate to the frontend directory:
   ```bash
   cd festivals-app/frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   ng serve
   ```

### Database

1. Create PostgreSQL database and setup connection string as an environment variable.
2. Migration code wi

ll create all tables, create `admin` account and fill in `countries` table.

## Usage

1. Access the app via the frontend URL (default: `http://localhost:4200`).
2. Register as an attendee, or login as `admin` (default password is defined in environment variables, by default it's `admin`) to register other administrators or organizers.
3. Enjoy!
