# Health Quest UI

A React-based user interface for the Health Quest application.

## Getting Started

These instructions will help you set up and run the project on your local machine for development and testing purposes.

### Prerequisites

- Node.js (v14 or later)
- npm (v6 or later)

### Installation

1. Clone the repository
2. Navigate to the project directory:
   ```
   cd frontend/health-quest-ui
   ```
3. Install dependencies:
   ```
   npm install
   ```

### Running the Application

To start the development server:

```
npm start
```

This will run the app in development mode. Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

### Building for Production

To build the app for production:

```
npm run build
```

This builds the app for production to the `build` folder. It correctly bundles React in production mode and optimizes the build for the best performance.

## Features

- OAuth2-based social login with Google
- Responsive design for both desktop and mobile

## OAuth Configuration

To use the Google login feature, you need to configure the OAuth credentials:

1. Create a project in the [Google Developer Console](https://console.developers.google.com/)
2. Enable the Google+ API
3. Create OAuth 2.0 credentials
4. Copy the `.env.example` file to a new file named `.env`:
   ```
   cp .env.example .env
   ```
5. Set your Google Client ID in the `.env` file:
   ```
   REACT_APP_GOOGLE_CLIENT_ID=your_actual_google_client_id
   ```

Note: The `.env` file is ignored by git to prevent sensitive information from being committed to the repository.

## Project Structure

```
health-quest-ui/
├── public/             # Public assets
├── src/                # Source files
│   ├── App.js          # Main application component
│   ├── App.css         # Application styles
│   ├── index.js        # Entry point
│   └── index.css       # Global styles
├── .env                # Environment variables (not committed to git)
├── .env.example        # Example environment variables template
└── package.json        # Project dependencies and scripts
```
