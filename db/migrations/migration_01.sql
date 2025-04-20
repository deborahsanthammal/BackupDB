-- Create the User table
CREATE TABLE User (
    UserID INT PRIMARY KEY AUTO_INCREMENT,
    Username VARCHAR(50) NOT NULL UNIQUE,
    PasswordHash VARCHAR(255) NOT NULL,
    Email VARCHAR(100) NOT NULL UNIQUE,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the UserProfile table
CREATE TABLE UserProfile (
    ProfileID INT PRIMARY KEY AUTO_INCREMENT,
    UserID INT NOT NULL,
    FirstName VARCHAR(50),
    LastName VARCHAR(50),
    DateOfBirth DATE,
    Address TEXT,
    PhoneNumber VARCHAR(15),
    FOREIGN KEY (UserID) REFERENCES User(UserID)
);

-- Create the UserSession table
CREATE TABLE UserSession (
    SessionID INT PRIMARY KEY AUTO_INCREMENT,
    UserID INT NOT NULL,
    SessionToken VARCHAR(255) NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ExpiresAt TIMESTAMP NOT NULL,
    FOREIGN KEY (UserID) REFERENCES User(UserID)
);

-- Create the UserAccount table
CREATE TABLE UserAccount (
    AccountID INT PRIMARY KEY AUTO_INCREMENT,
    UserID INT NOT NULL,
    AccountType VARCHAR(50) NOT NULL,
    AccountBalance DECIMAL(10, 2) DEFAULT 0.00,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (UserID) REFERENCES User(UserID)
);
