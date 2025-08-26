package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)

type model struct {
	students []string         // 学生名リスト
	cursor   int              // カーソル位置
	attend   map[int]struct{} // 出席済み集合（インデックスのみ保持）

	// 追加用の入力モード
	adding bool
	input  textinput.Model
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "新しい学生の名前を入力"
	ti.CharLimit = 64
	ti.Prompt = "> "

	return model{
		students: []string{"山田", "松本", "金城"},
		attend:   make(map[int]struct{}),
		adding:   false,
		input:    ti,
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		// 入力モード中のキー処理
		if m.adding {
			switch msg.String() {
			case "enter":
				name := strings.TrimSpace(m.input.Value())
				if name != "" {
					m.students = append(m.students, name)
					// 末尾に追加したのでカーソルを末尾へ
					m.cursor = len(m.students) - 1
				}
				m.adding = false
				m.input.Blur()
				m.input.SetValue("")
				return m, nil

			case "esc":
				// 追加をやめる
				m.adding = false
				m.input.Blur()
				m.input.SetValue("")
				return m, nil
			}

			// 文字入力は textinput に委譲
			var cmd tea.Cmd
			m.input, cmd = m.input.Update(msg)
			return m, cmd
		}

		// 通常モード中のキー処理
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.students)-1 {
				m.cursor++
			}

		case "enter", " ":
			if _, ok := m.attend[m.cursor]; ok {
				delete(m.attend, m.cursor)
			} else {
				m.attend[m.cursor] = struct{}{}
			}

		case "a":
			// 入力モードへ移行
			m.adding = true
			m.input.SetValue("")
			return m, m.input.Focus() // フォーカスしてカーソル点滅を開始
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.adding {
		// 入力モードの画面
		s := "出席管理 - 追加モード\n\n"
		s += "新しい学生名を入力して Enter。Esc でキャンセル。\n\n"
		s += m.input.View() + "\n"
		return s
	}

	// 通常の一覧画面
	s := "出席管理\n\n"
	for i, name := range m.students {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		check := " "
		if _, ok := m.attend[i]; ok {
			check = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, check, name)
	}
	s += "\n↑/k ↓/j 移動  space/enter 出欠切替  a 追加  q 終了\n"
	return s
}

func main() {
	if err := tea.NewProgram(initialModel(), tea.WithAltScreen()).Start(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

