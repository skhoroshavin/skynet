# SkyNET

This is a toy social network used for learning:
* golang & typescript
* hexagonal architecture
* handling high load

## Code structure

* backend
    * domain - core logic and interfaces
    * http - HTTP controller implementing REST API
    * mysql_storage - MySQL storage implementation
    * app - entry point, ties everything above together
* frontend
    * TODO