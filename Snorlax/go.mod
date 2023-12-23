module github.com/Chees3loaf/Pokedex/Snorlax

go 1.21.5

require github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/TMpocket v0.0.0

require (
	github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket v0.0.0 // indirect
	github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/Ballpocket v0.0.0 // indirect
	github.com/jung-kurt/gofpdf v1.16.2 // indirect
)

replace github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket => ../Ketchum/Sidepocket

replace github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/Ballpocket => ../Ketchum/Sidepocket/Ballpocket

replace github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/HMpocket => ../Ketchum/Sidepocket/HMpocket

replace github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/TMpocket => ../Ketchum/Sidepocket/TMpocket
