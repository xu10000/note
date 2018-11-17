syntax enable
set t_Co=256
set background=dark
set nocompatible              " be iMproved, required
set encoding=UTF-8
set backspace=2
set number
set splitbelow
" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
let g:airline#extensions#tabline#enabled = 1
let g:airline#extensions#tabline#tab_nr_type = 1
"let Tlist_Use_Left_Window=1
let g:tagbar_left=1
let Tlist_Show_One_File=1 
let Tlist_Exit_OnlyWindow=1
"设置airline箭头
let g:airline_powerline_fonts = 1
" 设置airline颜色
let g:airline_theme='luna'

nnoremap <Leader>b :bp<CR>
nnoremap <Leader>f :bn<CR>

" 查看buffers
nnoremap <Leader>l :ls<CR>

" 通过索引快速跳转
nnoremap <Leader>1 :1b<CR>
nnoremap <Leader>2 :2b<CR>
nnoremap <Leader>3 :3b<CR>
nnoremap <Leader>4 :4b<CR>
nnoremap <Leader>5 :5b<CR>
nnoremap <Leader>6 :6b<CR>
nnoremap <Leader>7 :7b<CR>
nnoremap <Leader>8 :8b<CR>
nnoremap <Leader>9 :9b<CR>
nnoremap <Leader>0 :10b<CR>
noremap <F3> :NERDTreeToggle<CR>
noremap <F2> :TagbarToggle<CR>
"保存文件
map <silent> <c-s> :update<CR> 
" 搜索快捷键
map <silent> <Space>sd :Ack<Space>
map <Space>sb :/
" 进入模式
nmap <Space>c :
" 关闭当前buffer
nmap <Space>bd :bdelete<CR>
" 打开终端
nmap <Space>' :term<CR>
filetype off                  " required
colorscheme gruvbox

call vundle#begin()

Plugin 'VundleVim/Vundle.vim'
Plugin 'ryanoasis/vim-devicons'
Plugin 'morhetz/gruvbox'
Plugin 'kien/ctrlp.vim'
Plugin 'scrooloose/nerdtree'
Plugin 'mileszs/ack.vim'
Plugin 'vim-airline/vim-airline'
Plugin 'vim-airline/vim-airline-themes'
Plugin 'majutsushi/tagbar'
Plugin 'fatih/vim-go'
Plugin 'Valloric/YouCompleteMe'

call vundle#end()            " required

filetype plugin indent on    " required

