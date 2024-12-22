# Car Rental Frontend

A modern React application for car rental services with blog and admin panel functionality.

## Features

- Car listing and details with reservation system
- Blog system with comments and likes
- User authentication and profile management
- Admin panel for managing cars, orders, users, and blog posts
- Responsive design with Tailwind CSS
- TypeScript for better type safety
- Modern state management with Redux Toolkit
- API integration with React Query

## Prerequisites

- Node.js 18.x or later
- npm 9.x or later

## Setup

1. Clone the repository
2. Install dependencies:
```bash
npm install
```

3. Create a `.env` file in the root directory with the following content:
```env
VITE_API_URL=http://localhost:5000/api
```

## Development

To start the development server:

```bash
npm run dev
```

The application will be available at `http://localhost:3000`.

## Build

To create a production build:

```bash
npm run build
```

## Linting

To run the linter:

```bash
npm run lint
```

## Project Structure

```
src/
├── assets/        # Static assets
├── components/    # Reusable components
├── hooks/         # Custom React hooks
├── pages/         # Page components
├── services/      # API services
├── store/         # Redux store and slices
├── types/         # TypeScript types and interfaces
└── utils/         # Utility functions
```

## Technologies Used

- React 18
- TypeScript
- Vite
- Redux Toolkit
- React Query
- React Router
- Tailwind CSS
- Axios
- Zod
- React Hook Form
