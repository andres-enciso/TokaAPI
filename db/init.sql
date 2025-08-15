USE TokaTasks;
GO

IF OBJECT_ID('dbo.Tasks', 'U') IS NULL
BEGIN
    CREATE TABLE dbo.Tasks (
        Id INT IDENTITY PRIMARY KEY,
        Titulo NVARCHAR(200) NOT NULL,
        Completado BIT NOT NULL DEFAULT(0),
        Fecha DATETIME2 NOT NULL DEFAULT SYSUTCDATETIME(),
    );

    INSERT INTO dbo.Tasks (Titulo, Completado) VALUES
    (N'Esta es la primer tarea', 0),
    (N'Segunda tarea', 1);
END
GO
IF OBJECT_ID('dbo.Users','U') IS NULL
BEGIN
    CREATE TABLE dbo.Users (
        ID INT IDENTITY PRIMARY KEY,
        Username NVARCHAR(100) NOT NULL UNIQUE,
        Password NVARCHAR(255) NOT NULL,
        CreatedAt DATETIME2 NOT NULL DEFAULT SYSUTCDATETIME(),
        UpdatedAt DATETIME2 NOT NULL DEFAULT SYSUTCDATETIME()
    );
END
GO