# Local Business Directory with Reviews (Backend)

This project is the backend for a Local Business Directory application that allows users to search for businesses, leave reviews, and explore local services. It is built in Go and uses Firebase for authentication and MongoDB for storing business and review data.

## Features

- **Business Listings**: Profiles with business details, photos, and contact information.
- **Review and Rating System**: Customers can leave reviews and ratings for businesses.
- **Search and Filter**: Users can search businesses by location, category, or service.
- **Promotional Features**: Highlight businesses with special offers or top ratings.

## Tech Stack

- **Go**: Backend language
- **Firebase**: Google Sign-In for user authentication
- **MongoDB**: Database for storing business listings and reviews

## Setup

1. Clone this repository:
    ```bash
    git clone https://github.com/yourusername/business-directory-backend.git
    cd business-directory-backend
    ```

2. Install Go dependencies:
    ```bash
    go mod tidy
    ```

3. Set up Firebase Authentication and MongoDB.

4. Configure environment variables for Firebase and MongoDB in your `.env` file.

5. Run the backend:
    ```bash
    go run main.go
    ```

## API Endpoints

- **POST /businesses**: Add a new business listing
- **GET /businesses**: Get a list of businesses (with optional filters)
- **POST /reviews**: Leave a review for a business
- **GET /reviews**: Get reviews for a business

## Future Improvements

- Add more filtering options (e.g., ratings, distance).
- Implement promotional features and business highlights.
