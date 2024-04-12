package guia5

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistrarAsistencia(t *testing.T) {
	asistencias := NewAsistencias(3, 5)
	asistencias.RegistrarAsistencia(0, 0)
	asistencias.RegistrarAsistencia(1, 0)
	asistencias.RegistrarAsistencia(1, 1)
	asistencias.RegistrarAsistencia(1, 2)
	asistencias.RegistrarAsistencia(1, 3)
	asistencias.RegistrarAsistencia(1, 4)
	asistencias.RegistrarAsistencia(2, 4)
	assert.True(t, asistencias.Asistio(0, 0))
	assert.True(t, asistencias.Asistio(1, 0))
	assert.True(t, asistencias.Asistio(1, 1))
	assert.True(t, asistencias.Asistio(1, 2))
	assert.True(t, asistencias.Asistio(1, 3))
	assert.True(t, asistencias.Asistio(1, 4))
	assert.True(t, asistencias.Asistio(2, 4))
	assert.False(t, asistencias.Asistio(0, 1))
	assert.False(t, asistencias.Asistio(0, 2))
	assert.False(t, asistencias.Asistio(0, 3))
	assert.False(t, asistencias.Asistio(0, 4))
	assert.False(t, asistencias.Asistio(2, 0))
	assert.False(t, asistencias.Asistio(2, 1))
	assert.False(t, asistencias.Asistio(2, 2))
	assert.False(t, asistencias.Asistio(2, 3))
}

func TestListarClase(t *testing.T) {
	asistencias := NewAsistencias(3, 5)
	asistencias.RegistrarAsistencia(0, 0)
	asistencias.RegistrarAsistencia(1, 0)
	asistencias.RegistrarAsistencia(1, 1)
	asistencias.RegistrarAsistencia(1, 2)
	asistencias.RegistrarAsistencia(1, 3)
	asistencias.RegistrarAsistencia(1, 4)
	asistencias.RegistrarAsistencia(2, 4)
	alumnos := asistencias.ListarClase(0)
	require.NotNil(t, alumnos)
	assert.Equal(t, 2, alumnos.Size())
	assert.True(t, alumnos.Contains(0))
	assert.True(t, alumnos.Contains(1))
	alumnos = asistencias.ListarClase(1)
	assert.Equal(t, 1, alumnos.Size())
	assert.True(t, alumnos.Contains(1))
	alumnos = asistencias.ListarClase(2)
	assert.Equal(t, 1, alumnos.Size())
	assert.True(t, alumnos.Contains(1))
	alumnos = asistencias.ListarClase(3)
	assert.Equal(t, 1, alumnos.Size())
	assert.True(t, alumnos.Contains(1))
	alumnos = asistencias.ListarClase(4)
	assert.Equal(t, 2, alumnos.Size())
	assert.True(t, alumnos.Contains(1))
	assert.True(t, alumnos.Contains(2))
}

func TestListarAlumno(t *testing.T) {
	asistencias := NewAsistencias(3, 5)
	asistencias.RegistrarAsistencia(0, 0)
	asistencias.RegistrarAsistencia(0, 1)
	asistencias.RegistrarAsistencia(0, 2)
	asistencias.RegistrarAsistencia(0, 3)
	asistencias.RegistrarAsistencia(0, 4)
	asistencias.RegistrarAsistencia(1, 0)
	asistencias.RegistrarAsistencia(1, 1)
	asistencias.RegistrarAsistencia(1, 2)
	asistencias.RegistrarAsistencia(1, 3)
	asistencias.RegistrarAsistencia(1, 4)
	asistencias.RegistrarAsistencia(2, 4)
	clases := asistencias.ListarAlumno(1)
	require.NotNil(t, clases)
	assert.Equal(t, 5, clases.Size())
	assert.True(t, clases.Contains(0))
	assert.True(t, clases.Contains(1))
	assert.True(t, clases.Contains(2))
	assert.True(t, clases.Contains(3))
	assert.True(t, clases.Contains(4))
	clases = asistencias.ListarAlumno(1)
	assert.Equal(t, 5, clases.Size())
	assert.True(t, clases.Contains(0))
	assert.True(t, clases.Contains(1))
	assert.True(t, clases.Contains(2))
	assert.True(t, clases.Contains(3))
	assert.True(t, clases.Contains(4))
	clases = asistencias.ListarAlumno(2)
	assert.Equal(t, 1, clases.Size())
	assert.True(t, clases.Contains(4))
}

func TestListarAsistencias(t *testing.T) {
	asistencias := NewAsistencias(3, 5)
	asistencias.RegistrarAsistencia(0, 0)
	asistencias.RegistrarAsistencia(0, 1)
	asistencias.RegistrarAsistencia(0, 2)
	asistencias.RegistrarAsistencia(0, 3)
	asistencias.RegistrarAsistencia(0, 4)
	asistencias.RegistrarAsistencia(1, 0)
	asistencias.RegistrarAsistencia(1, 1)
	asistencias.RegistrarAsistencia(1, 2)
	asistencias.RegistrarAsistencia(1, 3)
	asistencias.RegistrarAsistencia(1, 4)
	asistencias.RegistrarAsistencia(2, 2)
	clases := asistencias.ListarAsistencias()
	require.NotNil(t, clases)
	assert.Equal(t, 1, clases.Size())
	assert.True(t, clases.Contains(2))
}

func TestListarAsistenciaPerfecta(t *testing.T) {
	asistencias := NewAsistencias(3, 5)
	asistencias.RegistrarAsistencia(0, 0)
	asistencias.RegistrarAsistencia(0, 1)
	asistencias.RegistrarAsistencia(0, 2)
	asistencias.RegistrarAsistencia(0, 3)
	asistencias.RegistrarAsistencia(0, 4)
	asistencias.RegistrarAsistencia(1, 3)
	asistencias.RegistrarAsistencia(1, 4)
	asistencias.RegistrarAsistencia(2, 0)
	asistencias.RegistrarAsistencia(2, 1)
	asistencias.RegistrarAsistencia(2, 3)
	asistencias.RegistrarAsistencia(2, 4)
	alumnos := asistencias.ListarAsistenciaPerfecta()
	require.NotNil(t, alumnos)
	assert.Equal(t, 1, alumnos.Size())
	assert.True(t, alumnos.Contains(0))
}
