package client

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func Start() error {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
	return nil
}

// Add page states
type page int

const (
	splashPage page = iota
	loginPage
	chatPage
)

const gap = "\n" // Changed from "\n\n" to "\n"

// Modify model to include auth fields and current page
type model struct {
	currentPage page
	viewport    viewport.Model
	messages    []string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
	username    textinput.Model
	password    textinput.Model
	width       int
	height      int
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func initialModel() model {

	// get terminal width and height
	width, height, err := term.GetSize(0)
	if err != nil {
		width = 80  // fallback width
		height = 24 // fallback height
	}

	username := textinput.New()
	username.Placeholder = "Username"
	username.Focus()

	password := textinput.New()
	password.Placeholder = "Password"
	password.EchoMode = textinput.EchoPassword
	password.EchoCharacter = '•'

	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()
	ta.Prompt = "┃ "
	ta.CharLimit = 280
	ta.SetWidth(width)
	ta.SetHeight(2)
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetEnabled(false)

	vp := viewport.New(height, width)
	vp.SetContent(`Welcome to the chat room!
Type a message and press Enter to send.`)

	return model{
		currentPage: splashPage,
		textarea:    ta,
		messages:    []string{},
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		username:    username,
		password:    password,
		err:         nil,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		if m.currentPage == chatPage {
			m.viewport.Width = msg.Width
			m.textarea.SetWidth(msg.Width)
			m.viewport.Height = msg.Height - m.textarea.Height() - 1 // Adjusted padding
			if len(m.messages) > 0 {
				m.viewport.SetContent(lipgloss.NewStyle().Width(m.viewport.Width).Render(strings.Join(m.messages, "\n")))
			}
			m.viewport.GotoBottom()
		}

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			switch m.currentPage {
			case splashPage:
				m.currentPage = loginPage
				m.username.Focus()
				return m, nil
			case loginPage:
				if m.username.Focused() {
					m.username.Blur()
					m.password.Focus()
					return m, nil
				}
				if m.password.Focused() {
					// Here you would typically validate credentials
					m.currentPage = chatPage
					m.textarea.Focus()

					// Set up viewport dimensions for chat page
					m.viewport.Width = m.width
					m.viewport.Height = m.height - m.textarea.Height() - 1 // Adjusted padding

					// Initialize chat content
					m.viewport.SetContent(lipgloss.NewStyle().
						Width(m.viewport.Width).
						Render("Welcome to the chat room!\nType a message and press Enter to send."))

					return m, nil
				}
			case chatPage:
				if len(m.textarea.Value()) > 0 {
					newMessage := m.senderStyle.Render("You: ") + m.textarea.Value()
					m.messages = append(m.messages, newMessage)

					// Update viewport content with proper styling
					content := strings.Join(m.messages, "\n")
					styledContent := lipgloss.NewStyle().
						Width(m.viewport.Width).
						PaddingLeft(1).
						PaddingRight(1).
						Render(content)

					m.viewport.SetContent(styledContent)
					m.textarea.Reset()
					m.viewport.GotoBottom()
				}
			}
		}
	}

	// Handle input updates based on current page
	switch m.currentPage {
	case loginPage:
		if m.username.Focused() {
			var cmd tea.Cmd
			m.username, cmd = m.username.Update(msg)
			return m, cmd
		}
		if m.password.Focused() {
			var cmd tea.Cmd
			m.password, cmd = m.password.Update(msg)
			return m, cmd
		}
	case chatPage:
		var tiCmd, vpCmd tea.Cmd
		m.textarea, tiCmd = m.textarea.Update(msg)
		m.viewport, vpCmd = m.viewport.Update(msg)
		return m, tea.Batch(tiCmd, vpCmd)
	}

	return m, nil
}

func (m model) View() string {
	switch m.currentPage {
	case splashPage:
		return splashScreen()
	case loginPage:
		return loginScreen(m)
	case chatPage:
		return fmt.Sprintf(
			"%s%s%s",
			m.viewport.View(),
			gap,
			m.textarea.View(),
		)
	default:
		return "Loading..."
	}
}

func splashScreen() string {
	logo := `
 ____  _       _     _               
| __ )| | __ _| |__ | |__   ___ _ __ 
|  _ \| |/ _' | '_ \| '_ \ / _ \ '__|
| |_) | | (_| | |_) | |_) |  __/ |   
|____/|_|\__,_|_.__/|_.__/ \___|_|   
`
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Bold(true).
		Foreground(lipgloss.Color("5")).
		Render(logo + "\n\nPress Enter to continue...")
}

func loginScreen(m model) string {
	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Render(
			"Login\n\n" +
				m.username.View() + "\n\n" +
				m.password.View() + "\n\n" +
				"Press Enter to submit",
		)
}
