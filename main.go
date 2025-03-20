// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/ebfe/scard"
// )

// func main() {
// 	// Crear un contexto
// 	ctx, err := scard.EstablishContext()
// 	if err != nil {
// 		log.Fatalf("Error al establecer el contexto: %s", err)
// 	}
// 	defer ctx.Release()
// 	fmt.Println(ctx)

// 	// Obtener lista de lectores disponibles
// 	readers, err := ctx.ListReaders()
// 	if err != nil {
// 		log.Fatalf("Error al obtener la lista de lectores: %s", err)
// 	}
// 	if len(readers) == 0 {
// 		log.Fatal("No se encontraron lectores de tarjetas inteligentes")
// 	}

// 	// Seleccionar el primer lector de la lista
// 	reader := readers[0]

// 	// Conectar al lector
// 	card, err := ctx.Connect(reader, scard.ShareExclusive, scard.ProtocolAny)
// 	fmt.Println(card)
// 	fmt.Println(scard.AttrChannelId)
// 	fmt.Println(scard.AttrAsyncProtocolTypes)
// 	fmt.Println(scard.AttrAtrString)
// 	fmt.Println(scard.AttrVendorIfdSerialNo)
// 	if err != nil {
// 		log.Fatalf("Error al conectar al lector %s: %s", reader, err)
// 	}
// 	defer card.Disconnect(scard.LeaveCard)

// 	// Enviar comando APDU para leer datos de la tarjeta
// 	// El comando APDU específico puede variar dependiendo del tipo de tarjeta y la aplicación
// 	// apduCommand := []byte{0x3B, 0x00, 0x05, 0x00, 0x00} // SELECT MF
// 	// apduCommand := []byte{0x00, 0xA4, 0x04, 0x00, 0x0A, 0xD2, 0x76, 0x00, 0x00, 0x01, 0x24, 0x01, 0x00, 0x00} // Ejemplo: Leer los primeros 255 bytes
// 	// apduCommand := []byte{0x00, 0xB0, 0x00, 0x00, 0xFF} // Ejemplo: Leer los primeros 255 bytes
// 	apduCommand := []byte{0x00, 0xA4, 0x04, 0x00, 0x0A, 0xA0, 0x00, 0x00, 0x01, 0x67, 0x45, 0x4E, 0x02, 0x01, 0x01} // Ejemplo: Leer los primeros 255 bytes
// 	response, err := card.Transmit(apduCommand)

// 	if err != nil {
// 		log.Fatalf("Error al transmitir comando APDU: %s", err)
// 	}

// 	if response[len(response)-2] != 0x90 || response[len(response)-1] != 0x00 {
// 		log.Fatalf("Error en la respuesta APDU: %X", response)
// 	}

// 	// fmt.Println(response)
// 	fmt.Printf("Datos leídos de la tarjeta: %X\n", response)
// }

// package main

// import (
// 	"fmt"
// 	"github.com/ebfe/scard"
// 	"os"
// )

// func die(err error) {
// 	fmt.Println(err)
// 	os.Exit(1)
// }

// func waitUntilCardPresent(ctx *scard.Context, readers []string) (int, error) {
// 	rs := make([]scard.ReaderState, len(readers))
// 	for i := range rs {
// 		rs[i].Reader = readers[i]
// 		rs[i].CurrentState = scard.StateUnaware
// 	}

// 	for {
// 		for i := range rs {
// 			if rs[i].EventState&scard.StatePresent != 0 {
// 				return i, nil
// 			}
// 			rs[i].CurrentState = rs[i].EventState
// 		}
// 		err := ctx.GetStatusChange(rs, -1)
// 		if err != nil {
// 			return -1, err
// 		}
// 	}
// }

// func main() {

// 	// Establish a context
// 	ctx, err := scard.EstablishContext()
// 	if err != nil {
// 		die(err)
// 	}
// 	defer ctx.Release()

// 	// List available readers
// 	readers, err := ctx.ListReaders()
// 	if err != nil {
// 		die(err)
// 	}

// 	fmt.Printf("Found %d readers:\n", len(readers))
// 	for i, reader := range readers {
// 		fmt.Printf("[%d] %s\n", i, reader)
// 	}

// 	if len(readers) > 0 {

// 		fmt.Println("Waiting for a Card")
// 		index, err := waitUntilCardPresent(ctx, readers)
// 		if err != nil {
// 			die(err)
// 		}

// 		// Connect to card
// 		fmt.Println("Connecting to card in ", readers[index])
// 		card, err := ctx.Connect(readers[index], scard.ShareExclusive, scard.ProtocolAny)
// 		if err != nil {
// 			die(err)
// 		}
// 		defer card.Disconnect(scard.ResetCard)

// 		fmt.Println("Card status:")
// 		status, err := card.Status()
// 		if err != nil {
// 			die(err)
// 		}

// 		fmt.Printf("\treader: %s\n\tstate: %x\n\tactive protocol: %x\n\tatr: % x\n",
// 			status.Reader, status.State, status.ActiveProtocol, status.Atr)

// 		var cmd = []byte{0x00, 0xA4, 0x00, 0x00, 0x02, 0x3F, 0x00} // SELECT MF
// 		// var cmd = []byte{0x00, 0xa4, 0x00, 0x0c, 0x02, 0x3f, 0x00} // SELECT MF

// 		fmt.Println("Transmit:")
// 		fmt.Printf("\tc-apdu: % x\n", cmd)
// 		rsp, err := card.Transmit(cmd)
// 		if err != nil {
// 			die(err)
// 		}
// 		fmt.Printf("\tr-apdu: % x\n", rsp)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // Definir la función de tarea
// func task(n int) int {
// 	// Simula una tarea intensiva en cálculo
// 	// return n + n
// 	return n * n + 30*15 + 450 - 85 + (48*2 - 15)
// }

// // Worker function
// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	print(id)
// 	defer wg.Done()
// 	for n := range tasks {
// 		result := task(n)
// 		results <- result
// 	}
// }

// func main() {
// 	// Generar 200,000 números sobre los cuales ejecutar la tarea
// 	numbers := make([]int, 200_000)
// 	for i := 0; i < 200_000; i++ {
// 		numbers[i] = i + 1
// 	}

// 	// Capturar el tiempo de inicio
// 	startTime := time.Now()

// 	// Crear canales para las tareas y los resultados
// 	tasks := make(chan int, len(numbers))
// 	results := make(chan int, len(numbers))

// 	// Crear un wait group para esperar a que todas las goroutines terminen
// 	var wg sync.WaitGroup

// 	// Número de workers
// 	numWorkers := 4

// 	// Iniciar los workers
// 	for i := 0; i < numWorkers; i++ {
// 		wg.Add(1)
// 		go worker(i, tasks, results, &wg)
// 	}

// 	// Enviar tareas al pool de workers
// 	for _, num := range numbers {
// 		tasks <- num
// 	}
// 	close(tasks)

// 	// Esperar a que todos los workers terminen
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	// Recorrer y procesar los resultados a medida que se completen
// 	for result := range results {
// 		fmt.Printf("Resultado: %d\n", result)
// 	}

// 	// Capturar el tiempo de finalización
// 	endTime := time.Now()

// 	// Calcular y mostrar el tiempo total de ejecución
// 	elapsedTime := endTime.Sub(startTime)
// 	fmt.Printf("Tiempo total de ejecución golang: %.2f segundos\n", elapsedTime.Seconds())

// 	var pepe any = 16
// 	var pepe2 int = 16
// 	fmt.Println(pepe)
// 	fmt.Println(pepe2)
// }

// package main

// import (
//     "fmt"
//     "github.com/ebfe/scard"
//     "log"
// )

// func main() {
//     // Establece el contexto PC/SC
//     ctx, err := scard.EstablishContext()
//     if err != nil {
//         log.Fatalf("Error establishing context: %v", err)
//     }
//     defer ctx.Release()

//     // Lista los lectores disponibles
//     readers, err := ctx.ListReaders()
//     if err != nil {
//         log.Fatalf("Error listing readers: %v", err)
//     }

//     if len(readers) == 0 {
//         log.Fatalf("No readers found")
//     }

//     reader := readers[0]
//     fmt.Printf("Using reader: %s\n", reader)

//     // Conecta con la tarjeta
//     card, err := ctx.Connect(reader, scard.ShareShared, scard.ProtocolAny)
//     if err != nil {
//         log.Fatalf("Error connecting to card: %v", err)
//     }
//     defer card.Disconnect(scard.LeaveCard)

//     // Selecciona el applet del DNIe de RENIEC
//     selectAppCmd := []byte{0x00, 0xa4, 0x04, 0x00, 0x0A, 0xA0,0x00, 0x00, 0x00, 0x62, 0x03, 0x01, 0x0C, 0x06, 0x01}
//     // selectAppCmd := []byte{0x00, 0xA4, 0x04, 0x00, 0x10, 0xA0, 0x00, 0x00, 0x00, 0x77, 0x01, 0x00, 0x70, 0x0A, 0x10, 0x00, 0xF1, 0x00, 0x00, 0x01, 0x00}
//     response, err := transmitCommand(card, selectAppCmd)
//     if err != nil {
// 		log.Fatalf("Error transmitting select command: %v", err)
//     }
//     fmt.Printf("Select Applet Response: %x\n", response)

//     // Leer el número de DNI
//     readDNICmd := []byte{0x00, 0x00, 0x00, 0x00}
//     // readDNICmd := []byte{0x00, 0xB0, 0x00, 0x00, 0x00}
//     response, err = transmitCommand(card, readDNICmd)
//     if err != nil {
// 		log.Fatalf("Error reading DNI: %v", err)
//     }

// 	print(response)
//     fmt.Printf("DNI Number: %s\n", string(response))

//     // Aquí puedes añadir comandos adicionales para leer otros datos del DNIe.

// }

// // transmitCommand envía un comando APDU a la tarjeta y devuelve la respuesta
// func transmitCommand(card *scard.Card, cmd []byte) ([]byte, error) {
//     response, err := card.Transmit(cmd)
//     if err != nil {
//         return nil, err
//     }
//     return response, nil
// }

// package main

// import (
//     "fmt"
//     "github.com/ebfe/scard"
// )

// func main() {
//     // Establish a PC/SC context
//     context, err := scard.EstablishContext()
//     if err != nil {
//         fmt.Println("Error EstablishContext:", err)
//         return
//     }

//     // Release the PC/SC context (when needed)
//     defer context.Release()

//     // List available readers
//     readers, err := context.ListReaders()
//     if err != nil {
//         fmt.Println("Error ListReaders:", err)
//         return
//     }

//     // Use the first reader
//     reader := readers[0]
//     fmt.Println("Using reader:", reader)

//     // Connect to the card
//     card, err := context.Connect(reader, scard.ShareShared, scard.ProtocolAny)
//     if err != nil {
//         fmt.Println("Error Connect:", err)
//         return
//     }

//     // Disconnect (when needed)
//     defer card.Disconnect(scard.LeaveCard)

//     // Send select APDU
//     var cmd_select = []byte{0x00, 0xa4, 0x04, 0x00, 0x0A, 0xA0,
//   0x00, 0x00, 0x00, 0x62, 0x03, 0x01, 0x0C, 0x06, 0x01}
//     rsp, err := card.Transmit(cmd_select)
//     if err != nil {
//         fmt.Println("Error Transmit:", err)
//         return
//     }
//     fmt.Println(rsp)

//     // Send command APDU
//     var cmd_command = []byte{0x00, 0x00, 0x00, 0x00}
//     rsp, err = card.Transmit(cmd_command)
//     if err != nil {
//         fmt.Println("Error Transmit:", err)
//         return
//     }
//     fmt.Println(rsp)
// 	fmt.Printf("DNI Number: %s\n", string(rsp))
//     for i := 0; i < len(rsp)-2; i++ {
//         fmt.Printf("%c", rsp[i])
//     }
//     fmt.Println()
// }

// package main

// import (
//     "fmt"
//     "github.com/ebfe/scard"
//     "log"
// )

// func main() {
//     // Establece el contexto PC/SC
//     context, err := scard.EstablishContext()
//     if err != nil {
//         log.Fatalf("Error establishing context: %v", err)
//     }
//     defer context.Release()

//     // Lista los lectores disponibles
//     readers, err := context.ListReaders()
//     if err != nil {
//         log.Fatalf("Error listing readers: %v", err)
//     }

//     if len(readers) == 0 {
//         log.Fatalf("No readers found")
//     }

//     reader := readers[0]
//     fmt.Printf("Using reader: %s\n", reader)

//     // Conecta con la tarjeta
//     card, err := context.Connect(reader, scard.ShareShared, scard.ProtocolAny)
//     if err != nil {
//         log.Fatalf("Error connecting to card: %v", err)
//     }
//     defer card.Disconnect(scard.LeaveCard)

//     // Comando APDU para seleccionar el applet del DNIe de RENIEC
//     cmdSelect := []byte{0xA0, 0x00, 0x00, 0x00, 0x62, 0x03, 0x01, 0x0C, 0x06, 0x01}
//     rsp, err := card.Transmit(cmdSelect)
//     if err != nil {
//         log.Fatalf("Error transmitting select command: %v", err)
//     }
//     fmt.Printf("Select Applet Response: %x\n", rsp)

//     // Comando APDU para leer el número de DNI
//     cmdReadDNI := []byte{0x00, 0xA4, 0x04, 0x00}
//     rsp, err = card.Transmit(cmdReadDNI)
//     if err != nil {
//         log.Fatalf("Error reading DNI: %v", err)
//     }
//     fmt.Printf("DNI Number Response: %x\n", rsp)

//     // Imprimir el número de DNI
//     dniNumber := string(rsp[:len(rsp)-2]) // Eliminar los bytes SW1 y SW2
//     fmt.Printf("DNI Number: %s\n", dniNumber)

// }

// package main

// import (
// 	"bufio"
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"os/exec"
// 	"regexp"
// )

// func main() {
// 	// Ejecuta el comando certutil -scinfo
// 	cmd := exec.Command("certutil", "-scinfo")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatalf("Error executing certutil: %v", err)
// 	}

// 	// Procesa la salida para buscar el número de DNI
// 	scanner := bufio.NewScanner(&out)
// 	dniRegex := regexp.MustCompile(`[0-9]{8}`) // Asume que el DNI tiene 8 dígitos
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		// Busca patrones que coincidan con un DNI
// 		if dni := dniRegex.FindString(line); dni != "" {
// 			fmt.Printf("DNI found: %s\n", dni)
// 			return
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatalf("Error reading output: %v", err)
// 	}

// 	fmt.Println("DNI not found in certutil output")
// }

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})


	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}