package database

import "ARCTIC-WOLF/code/models"

//Currently Storing in Memory

// risks is our in-memory data store, keyed by the risk ID.
var MemoryOfRisks = make(map[string]models.Risk)
