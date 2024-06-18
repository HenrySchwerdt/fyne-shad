# Fyne Shad

This is a shadcn inspired component library for fyne. This lib is currently under development and has breaking changes. Please wait for v1.0.0 release before integrating it into your project.


# Components

## Badge

```go
badge := widgets.NewBadgeBuilder().
            Text("Badge").
            Build()
```

### Modifiers
- Background: Sets the background color to a defined color
- FontColor: Sets the fontcolor to a defined color
- FontSize: Sets the fontsize of the text
- FontWeight: Sets the fyne.TextStyle of the text