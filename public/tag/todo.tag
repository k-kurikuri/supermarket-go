<todo>

    <h3>{ opts.title }</h3>

    <ul>
        <li each={ items.filter(whatShow) }>
            <label class={ completed: done }>
                <input type="checkbox" checked={ done } onclick={ parent.toggle }> { title }
            </label>
        </li>
    </ul>

    <form onsubmit={ add }>
        <input ref="input" onkeyup={ edit }>
        <button disabled={ !text }>Add</button>

        <button type="button" disabled={ items.filter(onlyDone).length == 0 } onclick={ removeAllDone }>Del</button>
    </form>

    <!-- this script tag is optional -->
    <script>
        this.items = opts.items
        const self = this

        edit(e) {
            this.text = e.target.value
        }

        add(e) {
            e.preventDefault()
            if (this.text) {
                fetch('/add', {
                    method: 'POST',
                    headers: {
                      'content-type': 'application/json'
                    },
                    credentials: 'include',
                    body:JSON.stringify({title: this.text,})
                })
                        .then((res) => {
                            return res.json()
                        })
                        .then((json) => {
                          if (json.success) {
                            self.items.push({ title: self.text })
                            self.text = self.refs.input.value = ''
                            self.update()
                          }
                        })
            }
        }

        removeAllDone(e) {
            const delIndexes = []
            this.items.forEach((item, index) => {
              if(item.done) {
                delIndexes.push(index.toString())
              }
            })
            
            fetch('/delete', {
              method: 'DELETE',
              headers: {
                'content-type': 'application/json'
              },
              credentials: 'include',
              body:JSON.stringify({
                indexes: delIndexes,
              })
            }).then((res) => {
              return res.json
            })
            
            this.items = this.items.filter((item) => {
                return !item.done
            })
        }

        // an two example how to filter items on the list
        whatShow(item) {
            return !item.hidden
        }

        onlyDone(item) {
            return item.done
        }

        toggle(e) {
            const item = e.item
            item.done = !item.done

            fetch('/update', {
              method: 'PUT',
              headers: {
                'content-type': 'application/json'
              },
              credentials: 'include',
              body:JSON.stringify({
                index: this.items.indexOf(item).toString(),
                todo: {
                  title: item.title,
                  done: item.done,
                },
              })
            }).then((res) => {
              return res.json
            })

            return true
        }
    </script>

</todo>
