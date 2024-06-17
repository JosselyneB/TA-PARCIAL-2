package models

import (
    "database/sql"
    "log"
)

type Carro struct {
    ID            int
    Marca         string
    Modelo        string
    Color         string
    Placa         string
    PropietarioID int
}

// Método para crear un nuevo carro en la base de datos
func (c *Carro) CrearCarro(db *sql.DB) error {
    query := "INSERT INTO carros (marca, modelo, color, placa, propietario_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
    err := db.QueryRow(query, c.Marca, c.Modelo, c.Color, c.Placa, c.PropietarioID).Scan(&c.ID)
    if err != nil {
        log.Println("Error al insertar carro:", err)
        return err
    }
    return nil
}

// Método para obtener todos los carros de un cliente específico
func ObtenerCarrosPorCliente(db *sql.DB, propietarioID int) ([]Carro, error) {
    query := "SELECT id, marca, modelo, color, placa, propietario_id FROM carros WHERE propietario_id = $1"
    rows, err := db.Query(query, propietarioID)
    if err != nil {
        log.Println("Error al obtener carros:", err)
        return nil, err
    }
    defer rows.Close()

    var carros []Carro
    for rows.Next() {
        var c Carro
        err := rows.Scan(&c.ID, &c.Marca, &c.Modelo, &c.Color, &c.Placa, &c.PropietarioID)
        if err != nil {
            log.Println("Error al escanear carro:", err)
            return nil, err
        }
        carros = append(carros, c)
    }
    if err := rows.Err(); err != nil {
        log.Println("Error en filas de carros:", err)
        return nil, err
    }

    return carros, nil
}
