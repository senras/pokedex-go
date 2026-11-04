package pokecache

import (
	"testing"
	"time"
)

// Escenarios posibles
// 1. Agregar un elemento nuevo (clave no existente)
// 2. Agregar un elemento con una clave existente (sobrescribir valor)
// 3. Obtener un elemento existente (clave existente)
// 4. Obtener un elemento no existente (clave no existente)
// 5. Verificar que el valor recuperado sea el mismo que el agregado
func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(time.Minute)

	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	retrievedValue, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected to retrieve value for key %s, but got not found", key)
	}
	if string(retrievedValue) != string(value) {
		t.Errorf("Expected value %s for key %s, but got %s", value, key, retrievedValue)
	}
}

func TestCacheOverwrite(t *testing.T) {
	cache := NewCache(time.Minute)

	key := "testKey"
	value1 := []byte("testValue1")
	value2 := []byte("testValue2")

	cache.Add(key, value1)
	cache.Add(key, value2)

	retrievedValue, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected to retrieve value for key %s, but got not found", key)
	}
	if string(retrievedValue) != string(value2) {
		t.Errorf("Expected value %s for key %s, but got %s", value2, key, retrievedValue)
	}
}

func TestCacheGetNonExistent(t *testing.T) {
	cache := NewCache(time.Minute)

	key := "nonExistentKey"

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("Expected not to find value for key %s, but got found", key)
	}
}

func TestCacheExpiration(t *testing.T) {
	cache := NewCache(1 * time.Second)

	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	time.Sleep(2 * time.Second)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("Expected value for key %s to be expired, but got found", key)
	}
}